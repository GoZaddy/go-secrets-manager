# Secretly

Secretly is a golang package that allows you to store and retrieve secrets such as api keys on your local machine in an encrypted file.

# Basic Use

1. Open a vault (a file where your secrets are stored)
```
v, err := FileVault("test4.json", FileVaultOptions{CreateNew: true})
	if err != nil {
		log.Fatal(err)
	}
```
2. Store your secret
```
err = v.Set("password", "jghdbsh878yrjr98fni", "Gbola") //keyname, keyvalue, encryption key
	if err != nil {
		log.Fatal(err)
	}
```

3. Retrieve your secret
```
value, err := v.Get("password", "Gbola") //keyname, encryption key
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value) //jghdbsh878yrjr98fni
```

## Contributing
This package is open to contributions. Just create and issue and make a pull request!