/**
 * Created by Pulkit on 14-04-2016.
 */
function check(form) {
    if (form.UserId.value == "jerrygao" && form.password.value == "123456") {
        window.open('Dashboard.html')
        /*opens the target page while Id & password matches*/
    }
    else {
        alert("Error Password or Username")
        /*displays error message*/
    }
}
function validate(form){
    if(form.upassword.value != form.cpassword.value){
     alert("Passwords do not match");
    }
}
