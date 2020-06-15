# Airtable Convetor

Executable (écrit en [Go](https://golang.org/)) pour convertir les fichiers exportés depuis [Airtable](https://airtable.com/).

Par défaut, le script convertit les virgules (",") de sépration en point vigule de séparation (";"). Dans les textes longs, les caractères indésirables sont enlévés comme les caractères de saut de lignes et les points virgules. 

## Usage
### Command line

Le package s'utilise en ligne de commande, en aucune dépendance, de la manière qui suit :

--

En plaçant le script dans le même dossier que le fichier à convertir (fichier, ici, nommé "data.csv"), ouvre un terminal et lance une commande du style : 

```shell
> airtable-convertor data.csv
```
--

L'exécutable peut prendre en ligne de commande 4 arguments : 1 obligatoire et 3 optionnels

- ```filepath```: le chemin vers le fichier à convertir
- ```-eol ```: le caractère de saut de ligne (par défaut "\n")
- ```-sep ```: le caractère de séparation des données (par défaut ";")
- ```-uwc```: une liste (séparéé par des virgules) de caractère à enlever (en plus des caratèses de base)

### Interaction

Le package peut également être utilisé en mode interactif.

Si le chemin du fichier à traité n'est pas renseigné, il est alors demandé. Il suffit de le renseigner et d'appuyer sur entré pour lancer le processus.

## Téléchargement

Les versions précompilées sont téléchargeables dans la section "release" de la page ([ici](https://github.com/AmauryG13/airtable-convertor/releases)).

Il est nécessaire de bien sélectionner la version du script en fonction du système d'exploitation utilisé :
- Pour windows, téléchargez la version notée ```*-windows-*.exe```
- Pour macOs, téléchargez la version notée ```*-darwin-* ```
- Pour Linux, téléchargez la version notée ```*-linux-*```

Il est également de renommer le script selon son choix, sachant que l'appel du sript devra se faire avec le nom choisi.

### Compilation

Il est également possible de re - compiler l'exécutable à partir du code source à l'aide du logiciel ```cmake``` en utilisant le fichier __Makefile__.