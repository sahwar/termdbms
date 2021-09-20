package viewer

func GetHelpText() (help string) {
	help = `

##### Help:
	-p	database path (absolute)
    -d      specifies which database driver to use (sqlite/mysql)
    -a      enable ascii mode
	-h	prints this message
##### Controls:
###### MOUSE
	Scroll up + down to navigate table/text
	Move cursor to select cells for full screen viewing
###### KEYBOARD
	[WASD] to move around cells, and also move columns if close to edge
	[ENTER] to select selected cell for full screen view
	[UP/K and DOWN/J] to navigate schemas
    [LEFT/H and RIGHT/L] to navigate columns if there are more than the screen allows.
        Also to control the cursor of the text editor in edit mode
    [BACKSPACE] to delete text before cursor in edit mode
    [M(scroll up) and N(scroll down)] to scroll manually
	[Q or CTRL+C] to quit program
    [B] to toggle borders!
    [C] to expand column
    [P] in selection mode to write cell to file
    [R] to redo actions, if applicable
    [U] to undo actions, if applicable
	[ESC] to exit full screen view, or to enter edit mode
    [PGDOWN] to scroll down one views worth of rows
    [PGUP] to scroll up one views worth of rows
###### EDIT MODE (for quick, single line changes)
    [ESC] to enter edit mode with no pre-loaded text input from selection
    When a cell is selected, press [:] to enter edit mode with selection pre-loaded
    The text field in the header will be populated with the selected cells text. Modifications can be made freely
    [ESC] to clear text field in edit mode
    [ENTER] to save text. Anything besides one of the reserved strings below will overwrite the current cell
    [:q] to exit edit mode
    [:s] to save database to a new file (SQLite only)
    [:s!] to overwrite original database file (SQLite only). A confirmation dialog will be added soon
    [:h] to display help text
    [:new] opens current cell with a blank buffer
    [:edit] opens current cell in format mode
    [HOME] to set cursor to end of the text
    [END] to set cursor to the end of the text
###### FORMAT MODE (for editing lines of text)
    [ESC] to move between top control bar and format buffer
    [HOME] to set cursor to end of the text
    [END] to set cursor to the end of the text
    [:wq] to save changes and quit to main table view
    [:w] to save changes and remain in format view
    [:s] to serialize changes, non-destructive (SQLite only)
    [:s!] to serialize changes, overwriting original file (SQLite only)`

	return help
}
