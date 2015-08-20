package main

import "flag"
import "fmt"
import "io/ioutil"
import "github.com/google/go-github/github"
import "golang.org/x/oauth2"
import "os"
import "os/user"

func usage() {
    fmt.Println("Usage:")
    fmt.Println("\t gogist a.txt b.txt")
    fmt.Println("\t gogist -f filename.txt < a.txt")
    fmt.Println("\t pass -p to make the gist private")
    fmt.Println("\t pass -d 'the gist description' to add a description")
    os.Exit(1)
}

func getAccessToken() string {
    usr, err := user.Current()

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error checking current user.")
        os.Exit(1)
    }

    f := usr.HomeDir + "/.config/gogist/GH-ACCESS-TOKEN"

    bytes, err := ioutil.ReadFile(f)

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading file %s", f)
        os.Exit(1)
    }
    
    return string(bytes)   
}

func create(desc string, pub bool, files map[string]string) (*github.Gist, error) {

    ghat := getAccessToken()

     ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: ghat},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)

    client := github.NewClient(tc)

    f := make(map[github.GistFilename]github.GistFile)

    for k := range files {
        _k := github.GistFilename(k)
        f[_k] = github.GistFile{Content: github.String(files[k])}
    }

    gist := &github.Gist{
            Description: github.String(desc),
            Public: github.Bool(pub),
            Files: f,
    }

    gist, _, err := client.Gists.Create(gist)

    return gist, err
}

func main() {
    var fname = flag.String("f", "", "File name")
    var desc = flag.String("d", "", "Description")
    var public = flag.Bool("p", false, "Is public, default false")

    flag.Parse()

    files := flag.Args()

    fileMap := make(map[string]string)

    if len(files) > 0 {
        for _, f := range files {
            bytes, err := ioutil.ReadFile(f)
            if err == nil {
                fileMap[f] = string(bytes)
            } else {
                fmt.Fprintln(os.Stderr, "Error reading file %s", f)
                os.Exit(1)
            }
        }
    } else if *fname != "" {
        // TODO: Check whether there's anythin to read on stdin.
        bytes, err := ioutil.ReadAll(os.Stdin)

        if err == nil {
            fileMap[*fname] = string(bytes)
        }
    } else {
        usage()
    }

    res, err := create(*desc, *public, fileMap)
    if err == nil {
        fmt.Println(*res.HTMLURL)
    }
}
