use crate::domain::event::UsuarioCreado;

pub struct Usuario {
    pub usuario_id: String,
    pub nombre: String,
    pub email: String,
}

impl Usuario {
    pub fn from_event(event: UsuarioCreado) -> Self {
        Usuario {
            usuario_id: event.usuario_id,
            nombre: event.nombre,
            email: event.email,
        }
    }
}
