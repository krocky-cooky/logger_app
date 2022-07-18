var usernameElements = document.getElementsByClassName("username");
var usernameList = [];
for(var i = 0;i < usernameElements.length; ++i) {
    usernameList.push(usernameElements[i].innerHTML);
    console.log(usernameElements[i].innerHTML);
}

document.getElementById("username_input").addEventListener("input", function(e){
    var val = this.value;
    console.log(val);
    if(usernameList.includes(val)) {
        document.getElementById("username_alert").innerHTML = "そのユーザーネームは既に存在します";
    }
});