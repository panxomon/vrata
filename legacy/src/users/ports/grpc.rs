use crate::application::command_handler::CommandHandler;
use crate::application::query_handler::QueryHandler;
use std::io;

pub struct CliInterface {
    command_handler: CommandHandler,
    query_handler: QueryHandler,
}

impl CliInterface {
    pub fn new(command_handler: CommandHandler, query_handler: QueryHandler) -> Self {
        Self {
            command_handler,
            query_handler,
        }
    }

    pub fn start(&mut self) {
        println!("Bienvenido a la interfaz de línea de comandos. Ingrese comandos:");

        loop {
            let mut input = String::new();
            io::stdin().read_line(&mut input).expect("Error al leer la entrada");

            //TODO: Aquí debes analizar el input y dirigirlo al CommandHandler o QueryHandler correspondiente
            // Ejemplo:
            // if input.trim() == "crear_usuario" {
            //     let command = CrearUsuario { /* ... */ };
            //     self.command_handler.handle_command(Command::CrearUsuario(command));
            // }
        }
    }
}
