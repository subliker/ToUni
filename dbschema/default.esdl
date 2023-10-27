using extension graphql;
module default {
    type User {
        required username: str{
            constraint exclusive;
            constraint min_len_value(5);
        }
        required password: str{
            constraint min_len_value(8);
        }
        required created_at: datetime;
        updated_at: datetime;
        multi lessons: Lesson;
    }

    type Lesson {
        required owner: User;
        required name: str;
        description: str;
        required created_at: datetime;
        updated_at: datetime;
        data: str;
        files: array<str>;
    }
}
