# Kabarda Framework Core
- Clone the repository
- Navigate to directory
- Run make:build
- Copy the binary file to appropriate destination like so:
```
  cp dist/kabarda ~/your destination (please change the name if
  the distination is the same as curtrent directory)
```
- Run .YourbinaryFile new application_name:
```
    .YourbinaryFile new myapp
```
- Copy the binary file to your application root directory, so you can use
it to run your commands
- First navigate to your application and run: make restart
- Second generate auth(even if not needed) like so:  .YourbinaryFile make auth
- To see all commands run: .YourbinaryFile help