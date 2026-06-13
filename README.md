# 🟩 Analizador de Logs (JSON a CSV)

**Foco:** I/O eficiente, punteros vs valores, y metadatos explícitos mediante Struct Tags.

## 📋 Checklist de Requerimientos

* [ ] **[RF-2.1]** Recibir rutas de archivos mediante banderas de terminal: `--in` (JSON) y `--out` (CSV). Si el archivo no existe en la ruta actual (en el caso del out) hay que crearlo.
* [X] **[RF-2.2]** Mapear una estructura JSON con campos: `timestamp`, `level`, `message`, y `metadata`.
* [ ] **[RF-2.3]** Filtrar el flujo para exportar solo registros con niveles `ERROR` o `CRITICAL`.
* [ ] **[RF-2.4]** Capturar señales del sistema (`os/signal`) para cerrar archivos limpiamente si se interrumpe con `Ctrl+C`.

## 🛠️ Objetivos Técnicos (Conceptos a dominar)

* [X] Uso de go fmt, golint y go vet.
* [ ] Uso de `bufio.NewReader` y `bufio.NewWriter` para procesar archivos por streams (evitar carga total en RAM).
* [ ] Uso de la palabra clave `defer` para garantizar el cierre de recursos (`file.Close()`).
* [ ] Definición de *Struct Tags* nativos para el parseo de datos (`json:"level"`).
* [ ] Generación de documentación técnica local usando la herramienta nativa `go doc`.
* [ ] Uso básico de herramientas de prueba.

## 🏁 Definition of Done (DoD)

> Procesar un archivo de log simulado de 100MB en pocos milisegundos de forma secuencial y poder ver la documentación del paquete ejecutando `go doc` en la terminal.