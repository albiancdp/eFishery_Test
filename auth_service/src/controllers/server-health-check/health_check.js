const checkHealth = (req, res) => {
  const data = {
    uptime: process.uptime(),
    timestamp: new Date()
  };

  res.status(200).send(data);
};

export default checkHealth;
