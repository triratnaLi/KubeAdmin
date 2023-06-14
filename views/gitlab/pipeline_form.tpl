<!DOCTYPE html>
<html>
<head>
  <title>Pipeline Form</title>
</head>
<body>
<h1>Create Pipeline</h1>
<form action="/pipeline/create" method="post">
  <div>
    <label for="project_id">Project ID:</label>
    <input type="text" name="ProjectId" id="project_id">
  </div>
  <div>
    <label for="project_name">Project Name:</label>
    <input type="text" name="ProjectName" id="project_name">
  </div>
  <div>
    <label for="image_name">Image Name:</label>
    <input type="text" name="ImageName" id="image_name">
  </div>
  <div>
    <label for="app_name">App Name:</label>
    <input type="text" name="AppName" id="app_name">
  </div>
  <div>
    <label for="package_path">Package Path:</label>
    <input type="text" name="PackagePath" id="package_path">
  </div>
  <div>
    <label for="package">Package:</label>
    <input type="text" name="Package" id="package">
  </div>
  <div>
    <input type="submit" value="Create">
  </div>
</form>
</body>
</html>
