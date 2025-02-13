<html>
    <head>
        <title>Products</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">

    </head>
    <div class="container">
    <nav class="navbar navbar-expand-lg bg-body-tertiary  ">
        <div class="container-fluid">
          <a class="navbar-brand" href="#">Navbar scroll</a>
          <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarScroll" aria-controls="navbarScroll" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarScroll">
            <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
              <li class="nav-item">
                <a class="nav-link active" aria-current="page" href="#">Home</a>
              </li>
              
              
            </ul>
            <!-- <form class="d-flex" role="search"> -->
                <!-- <li class="nav-item"> -->
                    <a href="/logout" class="nav-link" aria-disabled="true">Logout</a>
                  <!-- </li> -->
            <!-- </form> -->
          </div>
        </div>
      </nav>
    </div>
    <body >
        <div class="container">
            <h1>Order Receipt</h1>
            <p>Thank you for your Order!</p>
            <h2>Order Summary</h2>
            <table class="table">
                <thead>
                  <tr>
                    <th scope="col">Product</th>
                    <th scope="col">Quantity</th>
                    <th scope="col">Price</th>
                    <th scope="col">Total</th>
                  </tr>
                </thead>
                <tbody>
                    {{range .cartItems}}
                  <tr>
                    <!-- <th scope="row">1</th> -->
                    <td>{{.ProductName}}</td>
                    <td>{{.Quantity}}</td>
                    <td>{{.Price}}</td>
                    <td>{{.Quantity | mul .Price}}</td>
                  </tr>
                  {{end}}
                  <tr>
                    <td colspan="3">Total</td>
                    <td>${{.order.Total}}</td>
                  </tr>
                </tbody>
              </table>
        </div>
    </body>
</html>