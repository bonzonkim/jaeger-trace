import initTrace from '../src/jaeger/root'
import * as express from 'express';
import apiRouter from '../src/app/router'

initTrace("test");

const PORT = 8080;
const app = express();
app.use('/api', apiRouter)


app.listen(PORT, () =>{
  console.log("server listening on port ", PORT)
})
