using crud.DB;
using crud.Models;
using crud.Views;
using GalaSoft.MvvmLight;
using GalaSoft.MvvmLight.Command;
using System.Collections.ObjectModel;
using System.Linq;
using System.Windows;

namespace crud.ViewModel
{
    public class MainViewModel : ViewModelBase
    {
        private LocalDB localDb;

        /// <summary>
        /// Initializes a new instance of the MainViewModel class.
        /// </summary>
        public MainViewModel()
        {
            localDb = new LocalDB();
            QueryCommand = new RelayCommand(Query);
            ResetCommand = new RelayCommand(() =>
            {
                Search = string.Empty;
                Query();
            });
            EditCommand = new RelayCommand<int>((t) => Edit(t));
            DeleteCommand = new RelayCommand<int>((t) => Delete(t));
            AddCommand = new RelayCommand(Add);
        }

        private ObservableCollection<Student> gridModelList;

        private string search = string.Empty;

        public string Search
        {
            get { return search; }
            set
            {
                search = value;
                RaisePropertyChanged();
            }
        }

        public ObservableCollection<Student> GridModelList
        {
            get { return gridModelList; }
            set
            {
                gridModelList = value;
                RaisePropertyChanged();
            }
        }

        public void Query()
        {
            var models = localDb.GetByName(Search);
            GridModelList = new ObservableCollection<Student>();

            if (models != null)
            {
                models.ForEach(arg =>
                {
                    GridModelList.Add(arg);
                });
            }
        }

        public void Add()
        {
            Student stu = new Student();
            UserView userView = new UserView(stu);
            var result = userView.ShowDialog();
            if (result.Value)
            {
                stu.Id = gridModelList.Max(x => x.Id) + 1;
                localDb.Add(stu);
                this.Query();
            }
        }

        public void Edit(int id)
        {
            var model = localDb.GetById(id);
            if (model != null)
            {
                UserView userView = new UserView(model);
                var result = userView.ShowDialog();

                if (result.Value)
                {
                    var newmodel = GridModelList.FirstOrDefault(x => x.Id.Equals(id));
                    if (newmodel != null)
                        newmodel.Name = model.Name;
                }
            }
        }

        public void Delete(int id)
        {
            var model = localDb.GetById(id);
            if (model != null)
            {
                var dr = MessageBox.Show($"{model.Name}ÇçÌèúÇµÇ‹Ç∑Ç©ÅH", "íÒé¶", MessageBoxButton.YesNo, MessageBoxImage.Question);
                if (dr == MessageBoxResult.Yes)
                {
                    localDb.Delete(id);
                    Query();
                }
            }
        }

        public RelayCommand QueryCommand { get; set; }
        public RelayCommand ResetCommand { get; set; }

        public RelayCommand<int> EditCommand { get; set; }
        public RelayCommand<int> DeleteCommand { get; set; }

        public RelayCommand AddCommand { get; set; }
    }
}