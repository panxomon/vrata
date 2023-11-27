
trait Command {}
trait CommandHandler {
    fn execute(&self, command: Box<dyn ICommand>);
}
struct CommandDispatcher {
    handlers: HashMap<String, Box<dyn ICommandHandler>>,
}

impl CommandDispatcher {
    fn dispatch(&self, command: Box<dyn ICommand>) {
        let command_type = format!("{:?}", command);
        if let Some(handler) = self.handlers.get(&command_type) {
            handler.execute(command);
        } else {
            println!("handler not found", command_type);
        }
    }


}