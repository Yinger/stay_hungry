using crud.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace crud.DB
{
    public class LocalDB
    {
        private List<Student> lsStudent;

        public LocalDB()
        {
            init();
        }

        private void init()
        {
            lsStudent = new List<Student>();
            for (int i = 0; i < 30; i++)
            {
                lsStudent.Add(new Student() { Id = i, Name = $"stu{i}" });
            }
        }

        public List<Student> Get()
        {
            return lsStudent;
        }

        public void Add(Student student)
        {
            lsStudent.Add(student);
        }

        public void Delete(int id)
        {
            var stu = lsStudent.FirstOrDefault(x => x.Id == id);
            if (stu != null)
                lsStudent.Remove(stu);
        }

        public List<Student> GetByName(string name)
        {
            return lsStudent.Where(x => x.Name.Contains(name)).ToList();
        }

        public Student GetById(int id)
        {
            Student stu = lsStudent.Where(x => x.Id.Equals(id)).FirstOrDefault();
            if (stu != null)
                return new Student() { Id = stu.Id, Name = stu.Name };
            else
                return null;
        }
    }
}