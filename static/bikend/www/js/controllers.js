angular.module('starter.controllers', [])

.controller('AppCtrl', function($scope, $ionicModal, $timeout) {

  // With the new view caching in Ionic, Controllers are only called
  // when they are recreated or on app start, instead of every page change.
  // To listen for when this page is active (for example, to refresh data),
  // listen for the $ionicView.enter event:
  //$scope.$on('$ionicView.enter', function(e) {
  //});

  // Form data for the login modal
  $scope.loginData = {};

  // Create the login modal that we will use later
  $ionicModal.fromTemplateUrl('templates/login.html', {
    scope: $scope
  }).then(function(modal) {
    $scope.modal = modal;
  });

  // Triggered in the login modal to close it
  $scope.closeLogin = function() {
    $scope.modal.hide();
  };

  // Open the login modal
  $scope.login = function() {
    $scope.modal.show();
  };

  // Perform the login action when the user submits the login form
  $scope.doLogin = function() {
    console.log('Doing login', $scope.loginData);

    // Simulate a login delay. Remove this and replace with your login
    // code if using a login system
    $timeout(function() {
      $scope.closeLogin();
    }, 1000);
  };
})

.controller('PlaylistsCtrl', function($scope) {
  $scope.playlists = [
    { title: 'Reggae', id: 1 },
    { title: 'Chill', id: 2 },
    { title: 'Dubstep', id: 3 },
    { title: 'Indie', id: 4 },
    { title: 'Rap', id: 5 },
    { title: 'Cowbell', id: 6 }
  ];
})

.controller('PlaylistCtrl', function($scope, $stateParams) {
        $scope.playlist = {
          user : {
            name: "Marty McFly",
            description: "Byke enthusiastic",
            image:"https://pbs.twimg.com/profile_images/1234618042/MartyMcfly.jpg",
          },
          parking : {
            address: "Copper box Arena",
            image: "https://upload.wikimedia.org/wikipedia/commons/b/bc/AlewifeBikeParking.agr.2001.JPG",
            limit: 3,
          },
          reviews :[
              {
                name: 'John Mayer',
                image: "http://www.technobuffalo.com/wp-content/uploads/2015/01/neytiri-avatar-5824.jpg",
                stars: 5,
                comment: 'lovely place to park bikes'
              },
              {
                name: 'Paul McFly',
                image: "http://images4.fanpop.com/image/photos/15200000/Avatar-Fan-Art-avatar-15271220-500-666.jpg",
                stars: 1,
                comment: 'The shower wasn\'t working, I had to leave al sweated to work!!!' 
              }
            ]
        };
})

.controller('MapCtrl', function($scope, $stateParams, esriLoader) {
    // initial map settings
        // initial map settings
        $scope.map = {
            center: {
                lng: -122.6819,
                lat: 45.5200
            },
            zoom: 13
        };

        // this example requires the extent module
        // so let's get that once the map is loaded
        $scope.mapLoaded = function(map) {
            esriLoader.require('esri/geometry/Extent').then(function(Extent) {
                // now that we have the Extent module, we can
                // wire up the click handler for bookmark buttons
                $scope.goToBookmark = function(bookmark) {
                    var extent = new Extent(bookmark.extent);
                    map.setExtent(extent);
                };
            });
        };
});
