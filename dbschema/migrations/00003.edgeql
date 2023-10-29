CREATE MIGRATION m1wd7wbldlc4baxd2pcg3hkvr7gbvtn5pzzjbrbzhabq7be4shorda
    ONTO m16lizutyrd4subsldnmjflopaq4rs62cpkrrk6e3oyn6e3he4zlcq
{
  ALTER TYPE default::User {
      CREATE REQUIRED PROPERTY role: std::str {
          SET REQUIRED USING (<std::str>{'User'});
          CREATE CONSTRAINT std::one_of('Admin', 'User');
      };
  };
};
