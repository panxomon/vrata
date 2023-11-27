use crate::domain::query::ObtenerUsuario;
use crate::domain::repository::UsuarioRepository;

pub struct QueryHandler {
    usuario_repository: Box<dyn UsuarioRepository>,
}

impl QueryHandler {
    pub fn new(usuario_repository: Box<dyn UsuarioRepository>) -> Self {
        Self { usuario_repository }
    }

    pub fn handle_query(&mut self, query: Query) {
        match query {
            Query::ObtenerUsuario(query) => self.obtener_usuario(query),
            //TODO: Agregar casos para   consultas
        }
    }

    fn obtener_usuario(&mut self, query: ObtenerUsuario) {
        //TODO:  manejar consulta ObtenerUsuario  interactuar con el repositorio y devolver resultados
    }
}
