suse crate::domain::command::{Command, CrearUsuario};
use crate::domain::repository::UsuarioRepository;

pub struct CommandHandler {
    usuario_repository: Box<dyn UsuarioRepository>,
}

impl CommandHandler {
    pub fn new(usuario_repository: Box<dyn UsuarioRepository>) -> Self {
        Self { usuario_repository }
    }

    pub fn handle_command(&mut self, command: Command) {
        match command {
            Command::CrearUsuario(command) => self.crear_usuario(command),
        }
    }

    fn crear_usuario(&mut self, command: CrearUsuario) {
        // TODO:interactuar con el repositorio y generar eventos
    }
}
