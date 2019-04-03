# Obfuscate

Encoding

```Go
encoded := obfuscate.Encode(myStructOrMap)
encoded := obfuscate.EncodeString("Hidden and safe")
```

Decoding

```Go
var decoded map[string]string
obfuscate.Decode(encodedString, &decoded)
```
