# NN (New Note)

`NN` is a tiny CLI tool that allows you to capture thoughts quickly. In short,
`nn` without arguments launches `$EDITOR` with a new file that has a timestamp
and a title row, in Markdown, in a designated inbox (by default `~/NewNotes/`)
The file's name is simply an incrementing number.

Alternatively, call `nn foo bar baz` and you get a file named `foo-bar-baz.md`,
and a filled-in title row. 

## Configuration

There's a config file at `~/.config/nn`, which is simple JSON, specifying the
inbox path and the current counter that is used for unnamed notes. 

The inbox path has to be an absolute path, becaue I didn't want to deal with
shell expansion. 

The counter is there so you can reset it programmatically if you process your
inboxes, if you want to. I enjoy having it rise indefinitely. 
