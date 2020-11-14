function getAll(entity) {
    fetch('https://faas-tarea6.netlify.app/api/' + entity)
      .then((response) => response.json())
        .then((data) => {
            fetch('/template/list/' + entity + '.html')
                .then((response) => response.text())
                .then((template) => {
                    var rendered = Mustache.render(template, data);
                    document.getElementById('content').innerHTML = rendered;
                });
        })
}

function getById(query, entity) {
    var params = new URLSearchParams(query);
    fetch('https://faas-tarea6.netlify.app/api/' + entity + '/?id=' + params.get('id'))
      .then((response) => response.json())
        .then((data) => {
            fetch('/template/detail/' + entity + '.html')
                .then((response) => response.text())
                .then((template) => {
                    var rendered = Mustache.render(template, data);
                    document.getElementById('content').innerHTML = rendered;
                });
        })
}

function home() {
    fetch('/template/home.html')
        .then((response) => response.text())
        .then((template) => {
            var rendered = Mustache.render(template, {});
            document.getElementById('content').innerHTML = rendered;
        });
}

function init() {
    router = new Navigo(null, false, '#!');
    router.on({
        '/contratos': function() {
            getAll('contratos');
        },
        '/equipos': function() {
            getAll('equipos');
        },
        '/jugadoras': function() {
            getAll('jugadoras');
        },
        '/contratoById': function(_, query) {
            getById(query, 'contratos');
        },
        '/equipoById': function(_, query) {
            getById(query, 'equipos');
        },
        '/jugadoraById': function(_, query) {
            getById(query, 'jugadoras');
        }
    });
    router.on(() => home());
    router.resolve();
}