'use strict';
var util = require('util'),
    path = require('path'),
    yeoman = require('yeoman-generator'),
    _ = require('lodash'),
    _s = require('underscore.string'),
    pluralize = require('pluralize'),
    asciify = require('asciify');

var GoGinMongodbGenerator = module.exports = function GoGinMongodbGenerator(args, options, config) {
  yeoman.generators.Base.apply(this, arguments);

  this.on('end', function () {
    this.installDependencies({ skipInstall: options['skip-install'] });
  });

  this.pkg = JSON.parse(this.readFileAsString(path.join(__dirname, '../package.json')));
};

util.inherits(GoGinMongodbGenerator, yeoman.generators.Base);

GoGinMongodbGenerator.prototype.askFor = function askFor() {

  var cb = this.async();

  console.log('\n' +
    '+-+-+ +-+-+-+ +-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
    '|g|o| |g|i|n| |m|o|n|g|o|d|b| |g|e|n|e|r|a|t|o|r|\n' +
    '+-+-+ +-+-+-+ +-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+\n' +
    '\n');

  var prompts = [{
    type: 'input',
    name: 'baseName',
    message: 'What is the name of your application?',
    default: 'myApp'
  }];

  this.prompt(prompts, function (props) {
    this.baseName = props.baseName;

    cb();
  }.bind(this));
};

GoGinMongodbGenerator.prototype.app = function app() {
  this.entities = [];
  this.resources = [];
  this.generatorConfig = {
    "baseName": this.baseName,
    "entities": this.entities,
    "resources": this.resources
  };
  this.generatorConfigStr = JSON.stringify(this.generatorConfig, null, '\t');

  this.template('_generator.json', 'generator.json');
  this.template('_package.json', 'package.json');
  this.copy('gitignore', '.gitignore');

  var serverDir = 'server/';
  var logsDir = serverDir + 'logs/';

  this.mkdir(serverDir);
  this.mkdir(logsDir);

  this.template('_main.go', 'main.go');
  this.template(serverDir + '_server.go', serverDir + 'server.go');
  this.template(serverDir + '_config.go', serverDir + 'config.go');
  this.template(logsDir + '_logger.go', logsDir + 'logger.go');
};

GoGinMongodbGenerator.prototype.projectfiles = function projectfiles() {

};
