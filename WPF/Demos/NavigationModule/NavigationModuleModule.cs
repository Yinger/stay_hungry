using NavigationModule.ViewModels;
using NavigationModule.Views;
using Prism.Ioc;
using Prism.Modularity;
using Prism.Regions;

namespace NavigationModule
{
    public class NavigationModuleModule : IModule
    {
        public void OnInitialized(IContainerProvider containerProvider)
        {
        }

        public void RegisterTypes(IContainerRegistry containerRegistry)
        {
            containerRegistry.RegisterForNavigation<ViewA>();
            containerRegistry.RegisterForNavigation<ViewB>();
            containerRegistry.RegisterDialog<ViewMsg, ViewMsgViewModel>();
        }
    }
}