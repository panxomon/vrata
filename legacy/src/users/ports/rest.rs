use crate::application::command_handler::CommandHandler;
use crate::application::query_handler::QueryHandler;
use warp::Filter;

pub struct RestInterface {
    command_handler: CommandHandler,
    query_handler: QueryHandler,
}

impl RestInterface {
    pub fn new(command_handler: CommandHandler, query_handler: QueryHandler) -> Self {
        Self {
            command_handler,
            query_handler,
        }
    }

    pub async fn start(&self) {
        //TODO: /*
        // let routes = /* Definir rutas aquí usando Warp o el marco que prefieras */;
        // warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
    }
}
