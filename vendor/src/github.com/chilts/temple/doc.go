// Copyright (c) 2016, Apps Attic Ltd (https://appsattic.com/) <chilts@appsattic.com>.

// Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby
// granted, provided that the above copyright notice and this permission notice appear in all copies.

// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN
// AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

/*

Package temple provides a simple way to read and execute a directory of templates. These templates can be kept in a cache when in production, or thrown away and re-read when in development. This saves you having to stop and start the server constantly.

Once in the cache, the templates are returned immediately when asked for.

    tmpl, err := temple.NewTemple("templates", "base.html", false)
    if err != nil {
        log.Fatal(err)
    }

    index, err := tmpl.Get("index.html")
    if err != nil {
        log.Fatal(err)
    }

    index.Execute(os.Stdout, nil)

*/
package temple
