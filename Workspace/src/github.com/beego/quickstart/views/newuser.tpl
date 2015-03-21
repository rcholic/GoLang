<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link href="/views/bower_components/bootstrap/dist/css/bootstrap.css" rel="stylesheet">
  <link href="/views/bower_components/bootstrap/dist/css/bootstrap-theme.css" rel="stylesheet">

</head>

<body>
  <div class="container">
    <div class="row">
      <h1>First Row</h1>
    </div>
    <div class="row-fluid">
    <div class="col-md-4">
    <div class="panel panel-info">
      <div class="panel panel-heading">
        <div class="panel-title">
          Panel Title
          </div>
        <div class="panel-body">
        <strong>left column</strong>
        <h1>Username: {{.username}} </h1>
        <h1>Password: {{.password}} </h1>
        </div>
      </div>
    </div>
    </div>
    <div class="col-md-8">
        <strong>right column</strong>
    </div>
    </div>
  </div>

  </body>
</body>
</html>
