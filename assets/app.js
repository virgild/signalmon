var css = require('./app.css');
var jquery = require('jquery');
var React = require('react');
var ReactDOM = require('react-dom');

var $ = jquery;

$(function(){
  ReactDOM.render(
    <h4>Hello, there signalmon!</h4>,
    document.getElementById('main')
  );
});
