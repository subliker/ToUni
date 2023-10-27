CREATE MIGRATION m1mk3kyyzjjcuiohc2btk7vjgxf2rymlqququ75yldo7x64yza75ra
    ONTO initial
{
  CREATE EXTENSION graphql VERSION '1.0';
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY created_at: std::datetime;
      CREATE REQUIRED PROPERTY password: std::str {
          CREATE CONSTRAINT std::min_len_value(8);
      };
      CREATE PROPERTY updated_at: std::datetime;
      CREATE REQUIRED PROPERTY username: std::str {
          CREATE CONSTRAINT std::exclusive;
          CREATE CONSTRAINT std::min_len_value(5);
      };
  };
  CREATE TYPE default::Lesson {
      CREATE REQUIRED LINK owner: default::User;
      CREATE REQUIRED PROPERTY created_at: std::datetime;
      CREATE PROPERTY data: std::str;
      CREATE PROPERTY description: std::str;
      CREATE PROPERTY files: array<std::str>;
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY updated_at: std::datetime;
  };
};
