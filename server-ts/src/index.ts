import '../env';
import logger from 'jet-logger';
import app from './server';
import envVars from '@shared/env-vars';


// Constants
const serverStartMsg = 'Koa Started on port localhost:',
    port = envVars.port;

// Start server
app.listen(port, () => {
    logger.imp(serverStartMsg + port);
});
