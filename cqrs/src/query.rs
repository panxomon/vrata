use std::marker::PhantomData;

trait Query {}

trait QueryHandler<TParameter, TResult> {
    fn retrieve(&self, query: TParameter) -> TResult;
}

trait QueryResult {}

struct QueryDispatcher<TContainer> {
    container: TContainer,
}

impl<TContainer> QueryDispatcher<TContainer> {
    fn new(container: TContainer) -> Self {
        Self { container }
    }

    fn dispatch<TParameter, TResult>(
        &self,
        query: TParameter,
    ) -> TResult
        where
            TParameter: Query,
            TResult: QueryResult,
    {
        let handler = self.container.resolve::<dyn QueryHandler<TParameter, TResult>>();
        handler.retrieve(query)
    }
}

struct UnityContainer;

impl UnityContainer {
    fn resolve<T>(&self) -> T {
        unimplemented!()
    }
}

fn main() {
    let container = UnityContainer;
    let query_dispatcher = QueryDispatcher::new(container);

    let query_result: QueryResult = query_dispatcher.dispatch::<_, ()>(/* provide query parameter here */);
}
