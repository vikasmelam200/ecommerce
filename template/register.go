<html>
    <head>
        <title>Register</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
        <script src="https://code.jquery.com/jquery-3.7.1.js" crossorigin="anonymous"></script>
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>

    </head>
    <body>
        <div class="container-fluid">
            <section class="vh-100">
                <div class="container-fluid h-custom">
                  <div class="row d-flex justify-content-center align-items-center h-100">
                    <div class="col-md-9 col-lg-6 col-xl-5">
                      <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
                        class="img-fluid" alt="Sample image">
                    </div>
                    <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
                      <!-- <form action="register" method="POST"> -->
                        <div class="d-flex flex-row align-items-center justify-content-center justify-content-lg-start">
                          <h1 >Register</h1>
                        </div>
                        <br><br>
                        <!-- Username input -->
                        <div data-mdb-input-init class="form-outline mb-4">
                            <input type="text" id="username" class="form-control form-control-lg"
                              placeholder="Enter a valid username" />
                            <label class="form-label" for="form3Example3">User Name</label>
                          </div>

                        <!-- Email input -->
                        <div data-mdb-input-init class="form-outline mb-4">
                          <input type="email" id="email" class="form-control form-control-lg"
                            placeholder="Enter a valid email address" />
                          <label class="form-label" for="form3Example3">Email address</label>
                        </div>
              
                        <!-- Password input -->
                        <div data-mdb-input-init class="form-outline mb-3">
                          <input type="password" id="password" class="form-control form-control-lg"
                            placeholder="Enter password" />
                          <label class="form-label" for="form3Example4">Password</label>
                        </div>
              
                        <div class="d-flex justify-content-between align-items-center">
                          <!-- Checkbox -->
                          <!-- <div class="form-check mb-0">
                            <input class="form-check-input me-2" type="checkbox" value="" id="form2Example3" />
                            <label class="form-check-label" for="form2Example3">
                              Remember me
                            </label>
                          </div> -->
                          <!-- <a href="#!" class="text-body">Forgot password?</a> -->
                        </div>
                        
                        <div class="text-center text-lg-start mt-4 pt-2">
                            <input type="submit" data-mdb-button-init data-mdb-ripple-init class="btn btn-primary btn-lg"
                            style="padding-left: 2.5rem; padding-right: 2.5rem;" id="singup" value="Sing Up">
                          <!-- <button  type="button" data-mdb-button-init data-mdb-ripple-init class="btn btn-primary btn-lg"
                            style="padding-left: 2.5rem; padding-right: 2.5rem;">Sing Up</button> -->
                          <p class="small fw-bold mt-2 pt-1 mb-0"> <a href="/login"
                              class="link-danger">Login</a></p>
                        </div>
              
                      <!-- </form> -->
                    </div>
                  </div>
                </div>
                <div class="toast align-items-center text-white bg-danger border-0" id="livetoast" role="alert" aria-live="assertive" aria-atomic="true">
                  <div class="toast-body" id="toast-message">
                   
                  </div>
                </div>
                <div
                  class="d-flex flex-column flex-md-row text-center text-md-start justify-content-between py-4 px-4 px-xl-5 bg-primary">
                 
                </div>
              </section>
        </div>
    </body>
    <script>
      $("#singup").click(function(){
        let username = $("#username").val().trim();
        let email = $("#email").val().trim();
        let password = $("#password").val().trim();

        if (username === ""){
          showToast("Username is required")
          return false
        }
        if (email === ""){
          showToast("email is required")
          return false
        }
        if (!validateEmail(email)){
          showToast("Please enter valid email id");
          return false;
        }
        if (password === ""){
          showToast("password is required")
          return false
        }

        console.log("username : "+username);
        console.log("email : "+email);
        console.log("password : "+password);
      })

      function showToast(message){
        $("#toast-message").text(message);
          let toast = new bootstrap.Toast(document.getElementById('livetoast'));
          toast.show()
      }
      function validateEmail(email){
        let re = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}/;
        return re.test(email);
      }
    </script>
</html>
