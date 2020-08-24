<p align="center">
    <a href="https://subkers.terisback.ru">
        <img alt="Subkers Logo" height="320" src="https://user-images.githubusercontent.com/26527529/91093603-15d30580-e673-11ea-835f-653a29d2d993.png">
    </a>
    <br>
    <a href="https://pkg.go.dev/github.com/Terisback/subkers?tab=doc">
        <img alt="go.dev" src="https://img.shields.io/badge/go.dev-007d9c?logo=go&logoColor=white&style=flat-square">
    </a>
</p>

<p align="center">
    <b>Subkers</b> is a tool for converting subtitles into markers for Adobe Audition.
    <br>
    Subkers works with any popular subtitle type. 
    <br> 
    Feel free to use it ^_^
</p>

## 👨🏻‍💻 Usage

<p align="center">
    <img alt="Web interface" src="https://user-images.githubusercontent.com/26527529/91093700-3ac77880-e673-11ea-8d15-d27384d1a12e.gif">
    <br>
    How to use web interface
    <br>
    <br>
    <img alt="To insert subtitles into Audition you need to press right button and choose import markers then choose your converted file" src="https://user-images.githubusercontent.com/26527529/91096835-1c17b080-e678-11ea-998b-d245bb18bb31.gif">
    <br>
    How to insert subtitle into Audition
    <br>
    <br>
    <img alt="To navigate markers you can use shortcut Ctrl+Alt+Arrows" src="https://user-images.githubusercontent.com/26527529/91097620-7ebd7c00-e679-11ea-95b2-fc61b56fa741.gif">
    <br>
    Marker navigation Ctrl+Alt+Arrows
    <br>
</p>



## ⚡️ CLI Usage

If you have installed Go you can install subkers with command
```bash
go install github.com/Terisback/subkers/cmd/subkers
```

Next you need to open console within subtitle files, and write this command
```bash
subkers "NAME-OF-YOUR-SUBTITLES"
```
Produces .csv file in same folder that you can import into your Adobe Audition.