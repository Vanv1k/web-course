document.addEventListener("DOMContentLoaded", function() {
    limitText();
});

function limitText() {
    var elements = document.getElementsByClassName("card-description");

    for (var i = 0; i < elements.length; i++) {
        var paragraph = elements[i];
        var maxLength = 200; // Максимальное количество символов
        var text = paragraph.textContent;
    
        if (text.length > maxLength) {
            paragraph.textContent = text.substring(0, maxLength) + "...";
        }
    }
}
