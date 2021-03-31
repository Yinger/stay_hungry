using Prism.Commands;
using Prism.Mvvm;
using Prism.Regions;
using Prism.Services.Dialogs;
using System;
using System.Windows;

namespace PrismModuleDemo.ViewModels
{
    public class MainWindowViewModel : BindableBase
    {
        private string _title = "Prism Application";
        private readonly IRegionManager _regionManager;
        private IRegionNavigationJournal _journal;
        private IDialogService _dialog;

        public string Title
        {
            get { return _title; }
            set { SetProperty(ref _title, value); }
        }

        public DelegateCommand<string> NavigateCommand { get; private set; }
        public DelegateCommand OpenACommand { get; private set; }
        public DelegateCommand OpenBCommand { get; private set; }
        public DelegateCommand GoBackCommand { get; private set; }

        public DelegateCommand GoForwardCommand { get; private set; }

        public MainWindowViewModel(IRegionManager regionManager, IDialogService dialog)
        {
            _regionManager = regionManager;
            _dialog = dialog;

            NavigateCommand = new DelegateCommand<string>(Navigate);
            OpenACommand = new DelegateCommand(openA);
            OpenBCommand = new DelegateCommand(openB);
            GoBackCommand = new DelegateCommand(goBack);
            GoForwardCommand = new DelegateCommand(goForward);
        }

        private void Navigate(string navigatePath)
        {
            if (navigatePath != null)
                _regionManager.RequestNavigate("ContentRegion", navigatePath, NavigationComplete);
        }

        private void NavigationComplete(NavigationResult result)
        {
            //System.Windows.MessageBox.Show(String.Format("Navigation to {0} complete. ", result.Context.Uri));
        }

        private void goBack()
        {
            _journal.GoBack();
        }

        private void goForward()
        {
            _journal.GoForward();
        }

        private void openB()
        {
            _regionManager.RequestNavigate("ContentRegion", "ViewB", arg =>
            {
                _journal = arg.Context.NavigationService.Journal;
            });
        }

        private void openA()
        {
            DialogParameters param = new DialogParameters();
            param.Add("Value", "Hello");
            _dialog.ShowDialog("ViewMsg", param, arg =>
            {
                if (arg.Result == ButtonResult.OK)
                {
                    var value = arg.Parameters.GetValue<string>("Value");
                    MessageBox.Show(value);
                }
                else
                {
                    MessageBox.Show("cancel");
                }
            });
            //NavigationParameters param = new NavigationParameters();
            //param.Add("Value", "Hello");
            //_regionManager.RequestNavigate("ContentRegion", "ViewA", param);
            _regionManager.RequestNavigate("ContentRegion", $"ViewA?Value=Hello", arg =>
            {
                _journal = arg.Context.NavigationService.Journal;
            });
        }
    }
}