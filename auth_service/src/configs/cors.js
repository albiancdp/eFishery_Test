const apiCors = {
  origin: ['*'],
  methods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
  allowedHeaders: ['Content-Type', 'Authorization', 'Accept', 'Content-Disposition', 'Allow-Control-Allow-Origin'],
  exposedHeaders: ['Content-Type', 'Authorization', 'Accept', 'Content-Disposition', 'Allow-Control-Allow-Origin'],
  maxAge: 3600,
  credentials: true,
  preflightContinue: false,
  optionsSuccessStatus: 200,
};

export default apiCors;
