var gulp = require("gulp")
var shell = require("gulp-shell")


gulp.task("install-binary", shell.task([
    'go install github.com/devignite/webapi'
]));

gulp.task("restart-supervisor",["install-binary"], shell.task([
    'supervisorctl restart webapi'
]));

gulp.task("watch", function (){
    gulp.watch('*' , ["install-binary","restart-supervisor"]);
});

gulp.task('default', ['watch'])