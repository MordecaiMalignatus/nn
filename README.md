# NN (New Note)

`NN` is a tiny CLI tool that allows you to capture thoughts quickly. In short,
`nn` without arguments launches `$EDITOR` with a new file that has a timestamp
and a title row, in Markdown, in a designated inbox (by default `~/NewNotes/`)
The file's name is simply an incrementing number if not changed.

If you do not change anything in the file opened, NN will delete the file to
avoid spamming your inbox with stuff you accidentally triggered.

### Naming files. 

Giving `NN` args will change the naming of your file from an incrementing
number to the args you gave it. Call `nn foo bar baz` and you get a file named
`foo-bar-baz.md`

Alternatively, you can edit the title line (The `# ` bit in the first line),
and it will get extracted and used no matter what you called `nn` with. 

### Pipes

As means to take a look at outputs later, NN supports pipes: `echo "hello!" |
nn test` will result in a file named `test.md` in your inbox, with contents
`hello!`. Using NN in a pipe does not generate a headline or date-stamp, just
the things you piped into it. 

### Configuration

There's a config file at `~/.config/nn`, which is simple JSON, specifying the
inbox path and the current counter that is used for unnamed notes. 

The inbox path has to be an absolute path, becaue I didn't want to deal with
shell expansion. 

The counter is there so you can reset it programmatically if you process your
inboxes, if you want to. I enjoy having it rise indefinitely. 
