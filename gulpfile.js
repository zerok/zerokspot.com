const Gulp = require('gulp');
const Sass = require('gulp-sass');
const rename = require('gulp-rename');
const merge = require('merge-stream');

Gulp.task('sass', ['copy:normalize'], () => {
    return Gulp.src('static/sass/main.scss')
        .pipe(Sass())
        .pipe(Gulp.dest('static/css'));
});

Gulp.task('copy:normalize', () => {
    return Gulp.src('node_modules/normalize.css/normalize.css')
        .pipe(rename({basename: '_normalize', extname: '.scss'}))
        .pipe(Gulp.dest('static/sass'));
});

Gulp.task('sass:watch', () => {
    Gulp.watch('static/sass/*.scss', ['sass']);
});

Gulp.task('default', ['sass', 'sass:watch']);
