# Gopher Go

This is a toy library for implementing gopher protocol servers in Go.

It's loosely structured after net/http and is api unstable.

# Usage

    s := gopher.Server{}

    h := gopher.HandleFunc(func(resp gopher.Response, req gopher.Request) {
       resp.WriteMenu(gopher.Menu{[]gopher.Link{
           gopher.Link{gopher.TextFile, "About", "/about", "localhost", 7070},
           gopher.Link{gopher.FullTextSearch, "Search", "/search", "localhost", 7070},
           }})
       resp.Write([]byte("Welcome to my Gopher hole!"))
       resp.End()
    }

    s.ListenAndServe("127.0.0.1:7070", h))
