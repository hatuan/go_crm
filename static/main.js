/**
 * Created by tuanha-01 on 5/6/2016.
 */

require.config({

    baseUrl: "",

    // alias libraries paths
    paths: {
        'application-configuration': 'app/application-configuration',
        'angular': 'bower_components/angular/angular',
        'angularAMD': 'bower_components/angularAMD/angularAMD',
        'jquery': 'bower_components/jquery/dist/jquery.min',
        'bootstrap': 'bower_components/bootstrap/dist/js/bootstrap.min',
        'blockUI': 'bower_components/angular-block-ui/dist/angular-block-ui.min',
        'ui.router': 'bower_components/angular-ui-router/release/angular-ui-router.min',
        'satellizer': 'bower_components/satellizer/satellizer',
        'pascalprecht.translate': 'bower_components/angular-translate/angular-translate.min',
        'toastr': 'bower_components/angular-toastr/dist/angular-toastr.tpls.min',
        'ajaxService': 'app/common/services/ajaxServices',
        'alertsService': 'app/common/services/alertsServices',
        'stateConfig': 'app/common/config/stateConfig'
    },
    // Add angular modules that does not support AMD out of the box, put it in a shim
    shim: {
        'angularAMD': ['angular'],
        'bootstrap': ['jquery'],
        'blockUI': ['angular'],
        'ui.router': ['angular'],
        'satellizer': ['angular'],
        'pascalprecht.translate': ['angular'],
        'toastr': ['angular']
    },

    // kick start application
    deps: ['application-configuration']
});
