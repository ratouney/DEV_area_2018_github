# Comment faire en sorte que ce projet fonctionne

Tu dois set la varialbe $GOPATH pour qu'elle inclue le repo.
Pas le dossier `/src`, le repo. SRC est un truc obligatoire pour que go fonctionne.

Les dépendances sont gérées de manière TRES carée et avec une structure très précise.

Tout code GO sur la machine (pas le projet, ca serait trop simple...) et localisé avec $GOPATH.

Donc d'abord, lancez un `go get` pour récuperer les dépendances de github et autres.

Ensuite `go run main.go` et ca devrait marcher.

Pour créer un segment de code ou un package, il faut faire un dossier au nom du package pour ensuite l'inclure dans les autres.

```
src/
    audio/
        audio.go
    video/
        video.go
    text/
        text.go
    main.go
```

Voici a quoi ca doit ressembler, mais attention !

Main.go peut importer ce qu'il veux, par contre `video` n'as pas accès a `audio`. Et l'inclusion circulaire ne marche pas. Voyons, quel langague stupide ferait ca. Les headers c'est pour les merdes...