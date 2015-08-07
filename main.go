package main

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/marcusolsson/passtray/pathtree"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

// newPassMenu returns a GTK menu of gpg files organized in the actual directory structure.
func newPassMenu(path string) *gtk.Menu {
	items := make(pathtree.Items)

	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".gpg" {
			entry := strings.TrimPrefix(strings.TrimSuffix(p, ".gpg"), path+"/")
			items.Add(strings.Split(entry, "/"))
		}
		return nil
	})

	return newMenuFromItems(items, "")
}

// newMenuFromItems returns a GTK menu from a pathtree.Items map. This function
// calls itself recursively and takes a relpath used to build the absolute path
// needed by the pass command.
func newMenuFromItems(items pathtree.Items, relpath string) *gtk.Menu {
	result := gtk.NewMenu()

	for k, v := range items {
		mi := gtk.NewMenuItemWithLabel(k)
		abspath := relpath + k

		if len(v) > 0 {
			mi.SetSubmenu(newMenuFromItems(v, abspath+"/"))
		} else {
			mi.Connect("activate", func() {
				exec.Command("pass", abspath, "-c").Run()
			})
		}

		result.Append(mi)
	}

	return result
}

// passPath returns the path to the default password storage.
func passPath() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}

	return filepath.Join(usr.HomeDir, ".password-store")
}

func main() {
	gtk.Init(&os.Args)

	glib.SetApplicationName("gtk-password-store")

	nm := newPassMenu(passPath())
	nm.ShowAll()

	si := gtk.NewStatusIconFromStock(gtk.STOCK_DIRECTORY)
	si.SetTitle("pass")
	si.SetTooltipMarkup("pass")
	si.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk.StatusIconPositionMenu, si, uint(cbx.Args(0)), uint32(cbx.Args(1)))
	})

	gtk.Main()
}
