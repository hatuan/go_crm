/**
 * Created by tuanha-01 on 5/11/2016.
 */
define(['angularAMD'], function (angularAMD) {

    var injectParams = ['$stateProvider', '$urlRouterProvider', '$authProvider'];
    stateConfig = function ($stateProvider, $urlRouterProvider, $authProvider) {

        $urlRouterProvider.otherwise("/");

        $stateProvider
            .state('login', {
                url: '/login',
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/user/login.html',
                        controller: 'LoginController',
                        controllerUrl: 'app/user/loginController',
                    })
                },
                resolve: {
                    skipIfAuthenticated: _skipIfAuthenticated
                }
            })
            .state('home', {
                url: '/',
                views: {
                    "master": angularAMD.route({
                        templateUrl: 'app/main/home.html',
                        controller: 'HomeController',
                        controllerUrl: 'app/main/homeController',
                    })
                },
                resolve: {
                    skipIfAuthenticated: _skipIfAuthenticated
                }
            });

        function _skipIfAuthenticated($q, $state, $auth) {
            var defer = $q.defer();
            if ($auth.isAuthenticated()) {
                defer.resolve(); // always return defer.resolve()
            } else {
                defer.resolve(); // always return defer.resolve()
            }
            return defer.promise;
        }

        function _redirectIfNotAuthenticated($q, $state, $auth) {
            var defer = $q.defer();
            if ($auth.isAuthenticated()) {
                defer.resolve(); // always return defer.resolve()
            } else {
                setTimeout(function () {
                    $state.go('login');
                }, 10);
                defer.resolve(); // always return defer.resolve()
            }
            return defer.promise;
        }
    }

    stateConfig.$inject = injectParams;

    return stateConfig;
});