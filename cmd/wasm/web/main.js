function fileFromInput() {
    var files = document.getElementById('files').files;

    for (var i = 0, f; f = files[i]; i++) {
        var reader = new FileReader();
        reader.onload = (function(theFile) {
            return function(e) {
                console.log(process(theFile.name, e.target.result));
            };
        })(f);
        reader.readAsText(f);
    }
}