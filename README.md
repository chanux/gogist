# gogist

*Command line github gist creator in go*

I really like github gist and use it a lot for sharing code/text. I like the convenience of command line. I like go. There.

## Installing

Copy the binary to somewhere on your $PATH

    cp gogist /usr/local/bin

Obtain Github API access token following the steps mentioned [here](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) (gogist just need access to gists so tick only gists when you create token)

gogist looks in `$HOME/.config/gogist/GH-ACCESS-TOKEN` file for the token. So put it there:

    mkdir $HOME/.config/gogist

    echo "YOUR ACCESS TOKEN" > $HOME/.config/gogist/GH-ACCESS-TOKEN

## Usage

To upload content of a.txt just:

    gogist a.txt

Upload multiple files (makes a single gist):

    gogist a.txt b.txt c.txt

    gogist *.txt

By default it reads from STDIN, and you can set a filename with `-f`:

    gogist -f test.txt < a.txt

gogist makes private gists by default. To make it public, pass `-p`:

    gogist -p a.txt

Use `-d` to add a description:

    gogist -d "I am the description" a.txt

## Other

[defunkt/gist](https://github.com/defunkt/gist) is a feature rich, cross platform cli gist creator.
