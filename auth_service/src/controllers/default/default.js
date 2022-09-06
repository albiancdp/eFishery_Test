const defaults = async (req, res) => {
  res.status(200).send({ message: 'Hello World!', description: 'Rest Api Running' });
};

export default defaults;
