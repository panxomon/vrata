use crate::domain::event::UsuarioCreado;

pub struct EventHandler;

impl EventHandler {
    pub fn handle_event(&mut self, event: Event) {
        match event {
            Event::UsuarioCreado(event) => self.usuario_creado(event),
            //TODO: Agregar casos para otros eventos
        }
    }

    fn usuario_creado(&mut self, event: UsuarioCreado) {
        //TODO:  manejar el evento UsuarioCreado + realizar acciones como enviar notificaciones, actualizar vistas, etc.
    }
}


