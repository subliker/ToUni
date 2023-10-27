CREATE MIGRATION m16lizutyrd4subsldnmjflopaq4rs62cpkrrk6e3oyn6e3he4zlcq
    ONTO m1mk3kyyzjjcuiohc2btk7vjgxf2rymlqququ75yldo7x64yza75ra
{
  ALTER TYPE default::User {
      CREATE MULTI LINK lessons: default::Lesson;
  };
};
