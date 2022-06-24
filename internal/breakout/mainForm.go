package breakout

import (
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/breakout/internal/game"
	"github.com/hultan/softteam/framework"
)

const applicationTitle = "breakout"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	AboutDialog *gtk.AboutDialog
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file and set up main window
	m.Window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)
	m.Window.SetApplication(app)
	m.Window.SetTitle("breakout main window")
	m.Window.SetSizeRequest(800, 600)
	m.Window.Connect("destroy", m.Window.Close)

	// Show the main window
	m.Window.ShowAll()

	// Start BreakOut game
	da := m.builder.GetObject("drawing_area").(*gtk.DrawingArea)
	g := game.NewBreakOut(m.Window, da)
	g.StartGame()
}
