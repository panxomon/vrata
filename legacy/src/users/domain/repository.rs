use crate::domain::aggregate::Usuario;
use crate::domain::event::UsuarioCreado;

pub trait UsuarioRepository {
    //TODO:  Método para guardar nuevo usuario en el repositorio
    fn guardar_usuario(&mut self, usuario: Usuario);
}

// Implementa el repositorio de usuarios usando un vector en memoria (puedes cambiar esto según tus necesidades)
pub struct InMemoryUsuarioRepository {
    usuarios: Vec<Usuario>,
}

impl InMemoryUsuarioRepository {
    // Constructor para crear una instancia del repositorio en memoria
    pub fn new() -> Self {
        InMemoryUsuarioRepository { usuarios: Vec::new() }
    }
}

// Implementa el trait UsuarioRepository para el repositorio en memoria
impl UsuarioRepository for InMemoryUsuarioRepository {
    fn guardar_usuario(&mut self, usuario: Usuario) {
        self.usuarios.push(usuario);
    }
}