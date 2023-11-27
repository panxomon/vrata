use std::time::SystemTime;

pub struct UsuarioCreado {
    pub usuario_id: String,
    pub nombre: String,
    pub email: String,
    pub fecha_creacion: SystemTime,
}

impl UsuarioCreado {
    pub fn new(usuario_id: String, nombre: String, email: String) -> Self {
        UsuarioCreado {
            usuario_id,
            nombre,
            email,
            fecha_creacion: SystemTime::now(),
        }
    }
}
