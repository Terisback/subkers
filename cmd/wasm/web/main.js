function fileFromInput() {
    var files = document.getElementById('file').files;

    for (var i = 0, f; f = files[i]; i++) {
        var reader = new FileReader();
        reader.onload = (function(theFile) {
            return function(e) {
                process(theFile.name, e.target.result);
            };
        })(f);
        reader.readAsText(f);
    }
}

function onDownload(name, text){ 
    var blob = new Blob([text], {type: "text/plain;charset=utf-8"});
    saveAs(blob, name);
}