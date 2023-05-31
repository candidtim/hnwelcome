# hnwelcome

     _________________________________________________________________
    / Show HN: `hnwelcome` is CLI app to show top Hacker News stories \
    | https://github.com/candidtim/hnwelcome                          |
    \ https://news.ycombinator.com/item?id=???                        /
     -----------------------------------------------------------------
            \   ^__^
             \  (oo)\_______
                (__)\       )\/\
                    ||----w |
                    ||     ||

A simple and fast CLI appliaction to show top
[Hacker News](https://news.ycombinator.com/) stories in the command line. Works
best with [`cowsay`](https://en.wikipedia.org/wiki/Cowsay). Use it like
[`fortune`](https://en.wikipedia.org/wiki/Fortune_(Unix)),
but tap into an endless supply of Hacker News in real time.

## Installation

With Go tool chain:

    $ go install github.com/candidtim/hnwelcome@latest

Or, download a binary for your OS from the
[Releases](https://github.com/candidtim/hnwelcome/releases) page.

If nothing else works, clone the repository and build with `go build .`.

## Usage

To see one of the top 5 stories:

    $ hnwelcome

    Hacking my “smart” toothbrush [467]
    https://kuenzi.dev/toothbrush/
    https://news.ycombinator.com/item?id=36128617

One of the newest stories:

    $ hnwelcome --newest

To choose randomly between the top 10 stories (instead of the default of top 5):

    $ hnwelcome -n 10

Pipe to `cowsay`:

    $ hnwelcome | cowsay -n
     _______________________________________________
    / Hacking my “smart” toothbrush [467]           \
    | https://kuenzi.dev/toothbrush/                |
    \ https://news.ycombinator.com/item?id=36128617 /
     -----------------------------------------------
            \   ^__^
             \  (oo)\_______
                (__)\       )\/\
                    ||----w |
                    ||     ||

See the built-in help for more options, or read below for customization.

### Show the top stories in new shell sessions

If you want to see one of the top stories every time you start a terminal, you
may run `hnwelcome | cowsay -n` in your shell's `.*rc` file
(`.bashrc`, `.zshrc`).

### Show the top stories in Vim on startup

For Vim users, if you want to see top HN stories in a welcome screen, my
recommendation is to use [vim-startify](https://github.com/mhinz/vim-startify)
with the following configuration:

    let g:startify_custom_header = split(system('hnwelcome | cowsay -n'), '\n')

### Customize the output format

To customize the output format, create a template file with Golang
[template syntax](https://pkg.go.dev/text/template) and point the `--template`
argument to it. See the
[Hacker News API documentation](https://github.com/HackerNews/API#items)
for the list of fields available for **stories**.

    $ hnwelcome --template PATH

For example, the default template is:

    {{title}} [{{score}}]
    {{url}}
    https://news.ycombinator.com/item?id={{id}}

### Latency and timeout

Due to the Internet connection and the API latency, `hnwelcome` may slow down
the shell or Vim startup time if you configure them to run it on startup. By
default, `hnwelcome` will try to fetch the story in under 1 second, or timeout
nicely otherwise. You can change the timeout value with a `--timeout` argument.
For example, if you have a somewhat slow Internet connection and you are ready
to wait for up to 2 seconds:

    hnwelcome --timetout 2000

## License

[The MIT License](http://opensource.org/licenses/MIT)
