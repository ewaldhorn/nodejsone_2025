# Node JS One 2025 Talk

Go and JavaScript Desktop Application Demo for my Node JS One 2025 talk.

An HTML frontend is built into a Go desktop application where Go provides the
app foundation and presents the UI via a web view component. This allows one to
access resources in a way that a browser can't, while still retaining the ability
to use the browser for the UI.

Repo: <https://github.com/ewaldhorn/nodejsone_2025>

## The problem

In an engineering environment product tests generate hundreds of thousands, if not
millions of data points from various sensors. Before uploading the results, engineers
want to verify that test parameters were executed as requested. Sometimes a part of
the product being tested failed early or did not engage, rendering the test mostly
useless. The desktop application allows them to monitor the local sensor cache database
in real-time, seeing on the graphs if or when a test becomes useless. They can then
either stop the test, reset it and run again or dig into the sensor data to figure
out what went wrong. All of this happens on their local machines where previously,
everything fed to the server and had to be discarded upon failure.

With the new workflow, feedback is instant and data only gets uploaded to the server
when the engineers have confirmed the test ran as planned. This saves a ton of time,
effort and network traffic.

## Our solution

There's no question that browsers have increased in power and functionality over
the last decade to the point where they rival native apps for building UI's. In fact,
a web-based UI is often easier to build and maintain than a native one! There are
so many fantastic libraries and frameworks available that one ignores the web at
your own peril.

That being said, while the web is a powerful platform, it does have some limitations
when it comes to data access and management. You don't really want your browser to
have unfettered access to your workstation, the risks are too great. Plus, there's a
limit to how much data you can sensibly load and manage in a browser session.

Of course, native solutions have limits as well, but, in general, have more access and
higher limits imposed on them. So this led to the exploration of using a web interface
backed by a native application by various companies. Things like Electron, Tauri etc.
all enable this functionality.

One other thing worth mentioning is that despite it having a web-based UI, we don't
have the same security concerns we'd have with a web-based app. Being able to process
files and even access local databases without needing internet or network access is
pretty handy for this use-case.

## Why Go though

Go is not something one naturally thinks of in these situations. What happened in this
instance is that there was existing Go API functionality and upcycling this code to
also serve in a desktop application made sense. So here we are, a desktop application
built in Go and JavaScript.

Some reasons I like Go:

- Compiled language with great concurrency support.
- Memory usage is generally low and the garbage collector is efficient.
- Compiles to a native binary which is small and easy to distribute.
- Being a native app, offline support is baked in.
- Full access to all system resources, it's a native app after all.

## Why HTML and JavaScript?

One could argue a fully native app would work, but the vast ecosystem that is the web
means that we have access to people with amazing UI skills. HTML-based interfaces are
very flexible and responsive and let's be honest, a tremendous amount of work and effort
have gone into making modern browsers insanely powerful.

Some reasons I like a web interface:

- Massive number of existing libraries, frameworks and talent.
- It's fast to iterate over prototyping, development and testing.
- Pretty consistent cross-platform support these days.
- Easy to update and adapt.
- UI experiments can happen without an app having been built yet.
- Front-end peeps do what they do with very few changes required.

## Ok but about Web Assembly

While there are technologies like Web Assembly that makes it possible to run Go
in the browser, it still means the code is running in the browser sandbox. With
this technique, we can use the browser for the UI and a standard application, with
all the access associated with these, as the supporting part of the solution.

## Functionality

The UI and Go components are able to interact via binding and this allows the UI
to call into the Go application and vice-versa. For example, the *Quit* button is
bound to a quit function that allows the UI to trigger an application exit.

We also bind this quit function to some short-cut keys, like you'd expect to find
in a regular desktop app to improve the user experience. This demonstrates how
existing JavaScript skills can still be used to create the overall UI experience
while tapping seamlessly into functionality provided by the Go component.

## Injected Scripts

The WebView component supports injecting JavaScript before the page load. In the
demo, it is used to inject the code that switches the button colours. In production
code, I've used to alter the UI by populating a template on the Go side and injecting
the JavaScript into the UI at runtime. This makes it possible for the desktop app
to pull feature flag data from a database, for example and update the UI accordingly,
just like you are to do in a normal web app.

## Process

1. The UI is built using standard HTML5, CSS3 and JavaScript.
2. It is then bundled into a single resource file, making it easy to embed.
3. A Go application is built, embedding the resource files.
4. The Go application enables two-way communication with the UI.
5. The webview handles UI interaction.
6. The Go app handles the rest.

## Technologies Used

- Go <https://go.dev/>
- WebView Go <https://github.com/webview/webview_go>
- JavaScript <https://developer.mozilla.org/en-US/docs/Web/JavaScript>
- NodeJS <https://nodejs.org/>
- WebPack <https://webpack.js.org/>

### Alternatives

A noteworthy alternative is the [Wails](https://wails.io/) project. I've used it
for some projects and it's great. For smaller applications, I've found WebView
to be simpler to maintain, but it's definitely worth knowing about Wails as well.
