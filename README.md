# Massive Attack Clash DB Website
The code within the repo is what comprises the Website

### How to deploy
Just simply build and run the executable
```bash
go build .
./main.exe
```
or 
```bash
go run .
```
(bottom option compiles and runs in one step)

### what it is reliant on
The program needs to have the database credentials stored properly within the project. all `.env` files have been gitignored and so are safe to be put in your local repository

### Progress
- [x] static webpage
- [x] go templates to inject information into a HTML page
- [ ] connecting to the database
- [ ] using HTMX to easily switch out contents in the HTML
- [ ] TailwindCSS for looks
- [ ] :momo:
- [ ] maybe a little JS/web Components for little effects
- [ ] go live
